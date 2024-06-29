import { Box, Button, MenuItem, TextField } from "@mui/material";
import React, { useState } from "react";
import { useRouter } from "next/navigation";

const server = process.env.NEXT_PUBLIC_API_URL || "http://127.0.0.1:9000";

interface InputFormProps {
  onListingCompleted?: () => void;
}

type formDataType = {
  image_name: string | File;
  text_name: string;
  class_name: string;
  price: number;
  sell_type: number;
};

export const InputForm: React.FC<InputFormProps> = (props) => {
  const { onListingCompleted } = props;
  const router = useRouter();
  const initialState = {
    image_name: "",
    text_name: "",
    class_name: "",
    price: 0,
    sell_type: 0,
  };
  const [values, setValues] = useState<formDataType>(initialState);

  const onValueChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setValues({
      ...values,
      [event.target.name]: event.target.value,
    });
  };
  const onFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setValues({
      ...values,
      [event.target.name]: event.target.files![0],
    });
  };
  const onSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    const data = new FormData();
    data.append("image", values.image_name);
    data.append("name", values.text_name);
    data.append("course_name", values.class_name);
    data.append("price", values.price.toString());
    data.append("sell_type", values.sell_type.toString());

    const token = localStorage.getItem("token");

    console.log("token:", token);

    fetch(server.concat("/items"), {
      //itemsのリンクで入力フォームが出るはず！！！
      method: "POST",
      mode: "cors",
      headers: {
        Authorization: `Bearer ${token}`,
      },
      body: data,
    })
      .then((response) => {
        console.log("POST status:", response.statusText);
        onListingCompleted && onListingCompleted();
        router.push("/home");
      })
      .catch((error) => {
        console.error("POST error:", error);
      });
  };

  const currencies = [
    {
      value: "rental",
      label: "レンタル",
    },
    {
      value: "sale",
      label: "売り出し",
    },
  ];

  return (
    <div className="Listing">
      <form onSubmit={onSubmit}>
        <Box mt={5}>
          <TextField
            type="file"
            // sx={{ display: "none" }}
            name="image_name"
            id="image_name"
            onChange={onFileChange}
            required
            fullWidth
            // label="写真"
            variant="outlined"
            inputProps={{
              multiple: true,
            }} //写真を複数入れられるかもしれないコード
            // variant="outlined"
          />
          <TextField
            type="text"
            name="text_name"
            id="text_name"
            onChange={onValueChange}
            required
            fullWidth
            label="教科書名"
            sx={{ mb: 3, mt: 3 }}
          />
          <TextField
            type="text"
            name="class_name"
            id="class_name"
            onChange={onValueChange}
            required
            fullWidth
            label="講義名"
            sx={{ mb: 3 }}
          />
          <TextField
            type="text"
            name="price"
            id="price"
            onChange={onValueChange}
            required
            fullWidth
            label="価格"
            sx={{ mb: 3 }}
          />
          <TextField
            select
            type="text"
            name="sell_type"
            id="sell_type"
            onChange={onValueChange}
            required
            fullWidth
            label="出品タイプ"
            helperText="選択してください"
            sx={{ mb: 3 }}
          >
            {currencies.map((option) => (
              <MenuItem key={option.value} value={option.value}>
                {option.label}
              </MenuItem>
            ))}
          </TextField>
          <Box textAlign="center">
            <Button type="submit" variant="outlined">
              List this item
            </Button>
          </Box>
        </Box>
      </form>
    </div>
  );
};

export default InputForm;
