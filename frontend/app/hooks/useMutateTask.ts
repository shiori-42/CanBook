// import axios from "axios";
// import { useQueryClient, useMutation } from "@tanstack/react-query";
// import { Task } from "../types";
// import useStore from "../store";
// import { useError } from "./useError";

// export const useMutateTask = () => {
//   const queryClient = useQueryClient();
//   const { switchErrorHandling } = useError();
//   const resetEditedTask = useStore((state) => state.resetEditedTask);

//   const createTaskMutation = useMutation(
//     (task: Omit<Task, "id" | "created_at" | "updated_at">) =>
//       axios.post<Task>(`${process.env.NEXT_PUBLIC_API_URLL}/tasks`, task),
//     {
//       onSuccess: (res) => {
//         const previousTasks = queryClient.getQueryData<Task[]>(["tasks"]);
//         if (previousTasks) {
//           queryClient.setQueryData(["tasks"], [...previousTasks, res.data]);
//         }
//         resetEditedTask();
//       },
//       onError: (err: any) => {
//         if (err.response.data.message) {
//           switchErrorHandling(err.response.data.message);
//         } else {
//           switchErrorHandling(err.response.data);
//         }
//       },
//     }
//   );
//   const updateTaskMutation = useMutation(
//     (task: Omit<Task, "created_at" | "updated_at">) =>
//       axios.put<Task>(`${process.env.NEXT_PUBLIC_API_URLL}/tasks/${task.id}`, {
//         name: task.name,
//       }),
//     {
//       onSuccess: (res, variables) => {
//         const previousTasks = queryClient.getQueryData<Task[]>(["tasks"]);
//         if (previousTasks) {
//           queryClient.setQueryData<Task[]>(
//             ["tasks"],
//             previousTasks.map((task) =>
//               task.id === variables.id ? res.data : task
//             )
//           );
//         }
//         resetEditedTask();
//       },
//       onError: (err: any) => {
//         if (err.response.data.message) {
//           switchErrorHandling(err.response.data.message);
//         } else {
//           switchErrorHandling(err.response.data);
//         }
//       },
//     }
//   );
//   const deleteTaskMutation = useMutation(
//     (id: number) =>
//       axios.delete(`${process.env.NEXT_PUBLIC_API_URLL}/tasks/${id}`),
//     {
//       onSuccess: (_, variables) => {
//         const previousTasks = queryClient.getQueryData<Task[]>(["tasks"]);
//         if (previousTasks) {
//           queryClient.setQueryData<Task[]>(
//             ["tasks"],
//             previousTasks.filter((task) => task.id !== variables)
//           );
//         }
//         resetEditedTask();
//       },
//       onError: (err: any) => {
//         if (err.response.data.message) {
//           switchErrorHandling(err.response.data.message);
//         } else {
//           switchErrorHandling(err.response.data);
//         }
//       },
//     }
//   );
//   return {
//     createTaskMutation,
//     updateTaskMutation,
//     deleteTaskMutation,
//   };
// };

//次の機能実装の際に参考にする
