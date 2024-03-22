// import React from "react";
// import { Login } from "../components/Login";

// const login = () => {
//   return <Login />;
// };

// export default login;
"use client";
import { useEffect } from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import axios from "axios";
import { Auth } from "../components/Auth";
import { Todo } from "../components/Todo";
// import { CsrfToken } from "../types";

function AppLogin() {
  useEffect(() => {
    // axios.defaults.withCredentials = true;
    // const getCsrfToken = async () => {
    //   const { data } = await axios.get<CsrfToken>(
    //     `${process.env.REACT_APP_API_URL}/csrf`
    //   );
    //   axios.defaults.headers.common["X-CSRF-Token"] = data.csrf_token;
    // };
    // getCsrfToken();
  }, []);
  return (
    // <BrowserRouter>
    //   <Routes>
    //     <Route path="/login" element={<Auth />} />
    //     <Route path="/aaa" element={<Todo />} />
    //   </Routes>
    // </BrowserRouter>
    <Auth />
  );
}

export default AppLogin;
