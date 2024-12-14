import { Box, Button, MenuItem, TextField, Typography } from "@mui/material";
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
  const [previewUrl, setPreviewUrl] = useState<string | null>(null); // プレビューURLを管理

  const onValueChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setValues({
      ...values,
      [event.target.name]: event.target.value,
    });
  };

  const onFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files![0];
    setValues({
      ...values,
      image_name: file,
    });

    // ファイルを読み込んでプレビューURLを設定
    const reader = new FileReader();
    reader.onload = () => {
      setPreviewUrl(reader.result as string); // 読み込んだ結果をプレビューURLに設定
    };
    if (file) {
      reader.readAsDataURL(file); // ファイルをData URLとして読み込む
    }
  };

  const onSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    const data = new FormData();
    data.append("image", values.image_name);
    data.append("text_name", values.text_name);
    data.append("class_name", values.class_name);
    data.append("price", values.price.toString());
    data.append("sell_type", values.sell_type.toString());

    const token = localStorage.getItem("token");

    console.log("token:", token);

    fetch(server.concat("/items"), {
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
    <Box className="Listing">
      <form onSubmit={onSubmit}>
        <Box mt={5}>
          {/* 画像プレビュー */}
          {previewUrl && (
            <Box mt={2} textAlign="center">
              {/* <Typography variant="body1">画像プレビュー:</Typography> */}
              <img
                src={previewUrl}
                alt="プレビュー"
                style={{
                  maxWidth: "50%",
                  height: "auto",
                  border: "1px solid #ddd",
                  borderRadius: "4px",
                  padding: "4px",
                }}
              />
              {/* next/Imageではエラーが出てしまったため急ぎの修正としてネイティブのimgタグを使用 */}
            </Box>
          )}
          {/* 画像アップロード */}
          <TextField
            type="file"
            name="image_name"
            id="image_name"
            onChange={onFileChange}
            required
            fullWidth
            variant="outlined"
            inputProps={{
              accept: "image/*", // 画像ファイルのみを許可
            }}
          />

          {/* 教科書名 */}
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
          {/* 講義名 */}
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
          {/* 価格 */}
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
          {/* 出品タイプ */}
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
          {/* 提出ボタン */}
          <Box textAlign="center">
            <Button type="submit" variant="outlined">
              List this item
            </Button>
          </Box>
        </Box>
      </form>
    </Box>
  );
};

export default InputForm;
