import { Box, Button, MenuItem, TextField, styled } from "@mui/material";
import React, { useEffect, useState } from "react";
import { useRouter } from "next/navigation";

const server = process.env.NEXT_PUBLIC_API_URL || "http://127.0.0.1:9000";

type EditProps = {
  itemId: string;
  onListingCompleted?: () => void;
  params: { id: string };
};

type formDataType = {
  image_name: string | File;
  text_name: string;
  class_name: string;
  price: number;
  sell_type: number;
};

const EditForm: React.FC<EditProps> = (props) => {
  const { itemId, onListingCompleted } = props;
  const router = useRouter();
  const initialState = {
    image_name: "",
    text_name: "",
    class_name: "",
    price: 0,
    sell_type: 0,
  };
  const [values, setValues] = useState<formDataType>(initialState);
  const [imageUrl, setImageUrl] = useState<string>("");

  useEffect(() => {
    const token = localStorage.getItem("token");
    fetch(`${server}/items/${itemId}`, {
      method: "GET",
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })
      .then((response) => response.json())
      .then((data) => {
        setValues({
          image_name: data.image_name,
          text_name: data.text_name,
          class_name: data.class_name,
          price: data.price,
          sell_type: data.sell_type,
        });
        setImageUrl(`${server}/images/${data.image_name}`);
      })
      .catch((error) => {
        console.error("GET error:", error);
      });
  }, [itemId]);

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
    setImageUrl(URL.createObjectURL(event.target.files![0])); //追加
  };

  const onSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();

    const data = new FormData();
    data.append("image", values.image_name);
    data.append("text_name", values.text_name);
    data.append("course_name", values.class_name);
    data.append("price", values.price.toString());
    data.append("sell_type", values.sell_type.toString());
    const token = localStorage.getItem("token");

    fetch(`${server}/items`, {
      method: "PUT",
      mode: "cors",
      headers: {
        Authorization: `Bearer ${token}`,
      },
      body: data,
    })
      .then((response) => {
        onListingCompleted && onListingCompleted();
        router.push("/home");
      })
      .catch((error) => {
        console.error("PUT error:", error);
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

  const StyledBox = styled(Box)(({ theme }) => ({
    position: "relative",
    overflow: "hidden",
    backgroundSize: "cover",
    backgroundPosition: "center",
    backgroundImage: `url(${imageUrl})`,
    width: "100%",
    height: 480,
    zIndex: 0,
    [theme.breakpoints.down("sm")]: {
      height: 400,
    },
  }));

  return (
    <Box>
      <form onSubmit={onSubmit}>
        <Box mt={5}>
          {imageUrl && <StyledBox />}
          <TextField
            type="file"
            name="image_name"
            id="image_name"
            onChange={onFileChange}
            required
            fullWidth
            variant="outlined"
            inputProps={{
              multiple: true,
            }} //写真を複数入れられるかもしれないコード
          />
          <TextField
            type="text"
            name="text_name"
            id="text_name"
            value={values.text_name}
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
            value={values.class_name}
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
            value={values.price}
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
            value={values.sell_type}
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
              Update this item
            </Button>
          </Box>
        </Box>
      </form>
    </Box>
  );
};

export default EditForm;
