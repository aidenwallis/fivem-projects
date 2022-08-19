import * as React from "react";

import { useSessionData } from "../auth";

export const CurrentUser: React.FC = () => {
  const session = useSessionData();
  return <h1>FiveM ID: {session?.identifiers.find((i) => i.startsWith("fivem:"))}</h1>;
};
