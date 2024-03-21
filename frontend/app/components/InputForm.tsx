import * as React from "react";
import { TextField, MenuItem, Typography, Box } from "@mui/material";

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

export default function BasicTextFields() {
  return (
    <Box
      component="form"
      sx={{
        width: "100%",
      }}
      noValidate
      autoComplete="off"
    >
      <TextField fullWidth label="教科書名" id="fullWidth" margin="normal" />
      <TextField fullWidth label="講義名" id="fullWidth" margin="normal" />

      <TextField
        fullWidth
        label="価格"
        id="outlined-number"
        type="number"
        // InputLabelProps={{
        //   shrink: true,
        // }}
        margin="normal"
      />

      <TextField
        fullWidth
        label="出品タイプ"
        id="outlined-select-currency"
        select
        helperText="選択してください"
        margin="normal"
      >
        {currencies.map((option) => (
          <MenuItem key={option.value} value={option.value}>
            {option.label}
          </MenuItem>
        ))}
      </TextField>

      <Typography
        mt={3}
        pt={1}
        borderBottom={1}
        borderColor={"#8c8c8c"}
        fontSize={15}
        fontWeight={"bold"}
      >
        商品の詳細
      </Typography>

      <TextField fullWidth label="fullWidth" id="fullWidth" margin="normal" />
    </Box>
  );
}
