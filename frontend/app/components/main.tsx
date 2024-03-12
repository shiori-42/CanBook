import * as React from "react";
import Container from "@mui/material/Container";

type TProps = {
  children: React.ReactNode;
};

const Main: React.FC<TProps> = ({ children }) => {
  return <Container maxWidth="sm">{children}</Container>;
};

export default Main;
