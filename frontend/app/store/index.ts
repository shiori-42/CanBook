// import { create } from "zustand";

// type EditedTask = {
//   id: number;
//   name: string;
// };

// type State = {
//   editedTask: EditedTask;
//   updateEditedTask: (payload: EditedTask) => void;
//   resetEditedTask: () => void;
// };

// const useStore = create<State>((set) => ({
//   editedTask: { id: 0, name: "" },
//   updateEditedTask: (payload) =>
//     set({
//       editedTask: payload,
//     }),
//   resetEditedTask: () => set({ editedTask: { id: 0, name: "" } }),
// }));

// export default useStore;

//次の機能実装の際に参考にする
