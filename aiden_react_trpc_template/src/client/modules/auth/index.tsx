import * as React from "react";

import { InferQueryOutput, setTRPCToken, TClientError, trpc } from "../utils/trpc";

export type Session = InferQueryOutput<"currentUser.get-session">;

export type AuthContextState = {
  error: TClientError | null;
  loading: boolean;
  session: Session | null;
};

export const AuthContext = React.createContext<AuthContextState>({
  error: null,
  loading: false,
  session: null,
});

export type TokenContextState = {
  currentToken: string | null;
};

export const TokenContext = React.createContext<TokenContextState>({
  currentToken: null,
});

export function useSession() {
  return React.useContext(AuthContext);
}

export function useSessionData() {
  const { session } = useSession();
  return session;
}

export function useSessionError() {
  const { error } = useSession();
  return error;
}

type Props = React.PropsWithChildren<{ blockUntilLoaded?: boolean }>;

const SessionProvider: React.FC<Props> = (props) => {
  const [currentToken, setCurrentToken] = React.useState<string | null>(null);

  React.useEffect(() => {
    const handler = (event: MessageEvent<{ type: "new-token"; token?: string }>) => {
      if (event.data?.type === "new-token" && event.data?.token) {
        setCurrentToken(event.data.token);
        setTRPCToken(event.data.token);
      }
    };

    window.addEventListener("message", handler);
    return () => {
      window.removeEventListener("message", handler);
    };
  }, [setCurrentToken]);

  if (props.blockUntilLoaded && !currentToken) {
    return null;
  }

  return <TokenContext.Provider value={{ currentToken }}>{props.children}</TokenContext.Provider>;
};

const InternalAuthProvider: React.FC<Props> = (props) => {
  const { isLoading, error, data } = trpc.useQuery(["currentUser.get-session"]);

  if (props.blockUntilLoaded && !data) {
    return null;
  }

  return (
    <AuthContext.Provider value={{ loading: isLoading, error, session: data || null }}>
      {props.children}
    </AuthContext.Provider>
  );
};

export const AuthProvider: React.FC<Props> = (props) => {
  return (
    <SessionProvider {...props}>
      <InternalAuthProvider {...props}>{props.children}</InternalAuthProvider>
    </SessionProvider>
  );
};
