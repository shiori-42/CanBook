// import { useEffect } from "react";
// import { BrowserRouter, Route, Routes } from "react-router-dom";
// import axios from "axios";
// import { CsrfToken } from "./types";
// import { Auth } from "../Auth";
// import { Todo } from "../Todo";

// function App() {
//   useEffect(() => {
//     axios.defaults.withCredentials = true;
//     const getCsrfToken = async () => {
//       const { data } = await axios.get<CsrfToken>(
//         `${process.env.NEXT_PUBLIC_API_URLL}/csrf`
//       );
//       axios.defaults.headers.common["X-CSRF-Token"] = data.csrf_token;
//     };
//     getCsrfToken();
//   }, []);
//   return (
//     <BrowserRouter>
//       <Routes>
//         <Route path="/" element={<Auth />} />
//         <Route path="/todo" element={<Todo />} />
//       </Routes>
//     </BrowserRouter>
//   );
// }

// export default App;
