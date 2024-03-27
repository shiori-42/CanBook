import React from "react";
import { Message } from "../chat/page";
import { Box, Typography, Paper, IconButton } from "@mui/material";
import { AccountCircle } from "@mui/icons-material";

const ChatBody = ({ data }: { data: Array<Message> }) => {
  return (
    <>
      {data.map((message: Message, index: number) => {
        if (message.type == "self") {
          return (
            <Box
              sx={{
                display: "flex",
                flexDirection: "column",
                mt: 2,
                width: "100%",
                alignItems: "flex-end",
              }}
              key={index}
            >
              <Typography variant="body2">{message.username}</Typography>
              <Paper
                sx={{
                  bgcolor: "primary.main",
                  color: "primary.contrastText",
                  px: 2,
                  py: 1,
                  mt: 1,
                  display: "inline-block",
                  borderRadius: "4px",
                }}
              >
                {message.content}
              </Paper>
            </Box>
          );
        } else {
          return (
            <Box
              sx={{
                mt: 2,
                width: "100%",
              }}
              key={index}
            >
              <Box display={"flex"} alignItems={"center"}>
                <IconButton sx={{ color: "#c6c6c6", p: 0, pr: 1 }}>
                  <AccountCircle fontSize="large" />
                </IconButton>
                <Typography variant="body2" sx={{ fontSize: "0.875rem" }}>
                  {message.username}
                </Typography>
              </Box>
              <Paper
                elevation={0} // 影をなくす
                sx={{
                  bgcolor: "grey.300", // 背景色をグレイに
                  color: "text.secondary", // テキスト色をダークセカンダリに
                  px: 2,
                  py: 1,
                  ml: 4,
                  mt: 1,
                  display: "inline-block",
                  borderRadius: "4px",
                }}
              >
                {message.content}
              </Paper>
            </Box>
          );
        }
      })}
    </>
  );
};

export default ChatBody;
