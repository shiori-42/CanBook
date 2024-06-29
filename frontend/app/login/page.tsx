// import React from "react";
// import { Login } from "../components/Login";

// const login = () => {
//   return <Login />;
// };

// export default login;
"use client";
import { useEffect } from "react";
import axios from "axios";
import { Auth } from "../components/Auth";
// import { CsrfToken } from "../types";

axios.defaults.withCredentials = true;
function AppLogin() {
  useEffect(() => {
    // const getCsrfToken = async () => {
    //   const { data } = await axios.get<CsrfToken>(
    //     `${process.env.NEXT_PUBLIC_API_URLL}/csrf`
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
