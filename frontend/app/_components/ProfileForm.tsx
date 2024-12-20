"use client";
import { Box, Button, TextField } from "@mui/material";

const ProfForm = () => {
  return (
    <>
      <Box
        width={"100%"}
        height={"100%"}
        borderColor={"divider"}
        mt={3}
        pb={3}
        display="flex"
        flexDirection="column"
        alignItems="center"
      >
        <TextField
          type="file"
          name="image_name"
          id="image_name"
          required
          fullWidth
          variant="outlined"
        />
      </Box>
      <Box>
        <TextField
          name="nickname"
          type="nickname"
          label="ニックネーム"
          variant="outlined"
          fullWidth
          required
          autoFocus
          // value={nickname}
        />
        <TextField
          name="college"
          type="college"
          label="大学名"
          variant="outlined"
          fullWidth
          required
          autoFocus
          // value={college}
          sx={{ mt: 3 }}
        />
        <TextField
          name="campus"
          type="campus"
          label="キャンパス名"
          variant="outlined"
          fullWidth
          required
          autoFocus
          // value={campus}
          sx={{ mt: 3 }}
        />
      </Box>
      <Box textAlign="center" mt={5}>
        <Button type="submit" variant="outlined" href="/profile">
          保存する
        </Button>
      </Box>
    </>
  );
};

export default ProfForm;
