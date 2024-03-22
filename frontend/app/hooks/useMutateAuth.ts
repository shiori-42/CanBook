import axios from "axios";
import { useNavigate } from "react-router-dom";
import { useMutation } from "@tanstack/react-query";
import useStore from "../store";
import { Credential } from "../types";
import { useError } from "../hooks/useError";
import { useRouter } from "next/navigation";

export const useMutateAuth = () => {
  const router = useRouter();
  const resetEditedTask = useStore((state) => state.resetEditedTask);
  const { switchErrorHandling } = useError();
  const loginMutation = useMutation(
    async (user: Credential) =>{
      console.log(process.env.NEXT_PUBLIC_API_URL)
      await axios.post(`${process.env.NEXT_PUBLIC_API_URL}/login`, user) //ここが怪しい「/login」が正しいのか？バックエンド側で/loginではないものを書いてほしいかも
    },
    {
      onSuccess: () => {
        router.push("/todo");
      },
      onError: (err: any) => {
        if (err.response.data.message) {
          switchErrorHandling(err.response.data.message);
        } else {
          switchErrorHandling(err.response.data);
        }
      },
    }
  );
  const registerMutation = useMutation(
    async (user: Credential) =>
      await axios.post(`${process.env.NEXT_PUBLIC_API_URL}/signup`, user),
    {
      onError: (err: any) => {
        if (err.response.data.message) {
          switchErrorHandling(err.response.data.message);
        } else {
          switchErrorHandling(err.response.data);
        }
      },
    }
  );
  const logoutMutation = useMutation(
    async () => await axios.post(`${process.env.NEXT_PUBLIC_API_URL}/logout`),
    {
      onSuccess: () => {
        resetEditedTask();
        router.push("/login");
      },
      onError: (err: any) => {
        if (err.response.data.message) {
          switchErrorHandling(err.response.data.message);
        } else {
          switchErrorHandling(err.response.data);
        }
      },
    }
  );
  return { loginMutation, registerMutation, logoutMutation };
};
