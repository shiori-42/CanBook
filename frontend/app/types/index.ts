export type Task = {
  id: number;
  name: string;
  created_at: Date;
  updated_at: Date;
};

// export type CsrfToken = {
//   csrf_token: string;
// };

export type Credential = {
  name: string;
  email: string;
  password: string;
};
