import axios from "axios";
import { useMutation } from "@tanstack/react-query";
// import useStore from "../store";
import { useError } from "../hooks/useError";
import { useRouter } from "next/navigation";

export const useMutateAuth = () => {
  const router = useRouter();
  // const resetEditedTask = useStore((state) => state.resetEditedTask);
  const { switchErrorHandling } = useError();
  const loginMutation = useMutation(
    async (user: { email: string; password: string }) => {
      console.log(process.env.NEXT_PUBLIC_API_URL);

      const res = await axios.post(
        `${process.env.NEXT_PUBLIC_API_URL}/login`,
        user
      );

      return res.data;
    },
    {
      onSuccess: (data) => {
        localStorage.setItem("token", data.token);

        router.push("/home");
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
    async (user: {
      name: string;
      email: string;
      password: string;
      college: string;
      campus: string;
    }) => await axios.post(`${process.env.NEXT_PUBLIC_API_URL}/signup`, user),
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
        localStorage.removeItem("token"); // トークンをローカルストレージから削除

        // resetEditedTask();
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
