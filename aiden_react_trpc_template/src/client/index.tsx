import "./styles/index.scss";

import * as React from "react";
import { createRoot } from "react-dom/client";

import { AuthProvider } from "./modules/auth";
import { CurrentUser } from "./modules/current-user";
import { TRPCProvider } from "./modules/utils/trpc";

const App: React.FC = () => {
  return (
    <TRPCProvider>
      <AuthProvider blockUntilLoaded>
        <CurrentUser />
      </AuthProvider>
    </TRPCProvider>
  );
};

const mount = document.getElementById("app-mount");
mount && createRoot(mount).render(<App />);
