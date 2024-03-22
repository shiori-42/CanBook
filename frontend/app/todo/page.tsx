"use client";
import { useEffect } from "react";
import { Todo } from "../components/Todo";

function AppTodo() {
  useEffect(() => {}, []);
  return <Todo />; //これを出品画面にする
}

export default AppTodo;
