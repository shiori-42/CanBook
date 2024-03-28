"use client";
import React, { useState, useEffect } from "react";
import { Box, TextField, Button } from "@mui/material";
import ChatBody from "../components/ChatBody";

export type Message = {
  content: string;
  senderID: string;
  recipientID: string;
  type: "recv" | "self";
  username: string;
};

const ChatPage = ({ params }: { params: { itemId: string } }) => {
  const [messages, setMessages] = useState<Array<Message>>([]);
  const [inputMessage, setInputMessage] = useState("");
  const [ws, setWs] = useState<WebSocket | null>(null);
  const itemId = params.itemId;

  useEffect(() => {
    const token = localStorage.getItem("token");
    if (token && itemId) {
      const wsProtocol = window.location.protocol === "https:" ? "wss:" : "ws:";
      const wsUrl = `${wsProtocol}://${process.env.NEXT_PUBLIC_API_URL?.replace(/^https?:\/\//, '')}/ws/${itemId}`;
      const newWs = new WebSocket(wsUrl);

      newWs.onopen = () => {
        console.log("WebSocket connected");
        newWs.send(JSON.stringify({ token }));
      };
  
      newWs.onmessage = (event) => {
        const message = JSON.parse(event.data) as Message;
        setMessages((prevMessages) => [...prevMessages, message]);
      };
  
      newWs.onclose = () => {
        console.log("WebSocket disconnected");
      };
  
      newWs.onerror = (error) => {
        console.error("WebSocket error:", error);
      };
  
      setWs(newWs);
  
      return () => {
        newWs.close();
      };
    }
  }, [itemId]);

  const sendMessage = () => {
    console.log("sendMessageが呼び出されました"); // デバッグログ
    if (!ws) {
        console.log("WebSocketがnullです");
        return;
    }
    if (ws.readyState !== WebSocket.OPEN) {
        console.log(`WebSocketの状態がOPENではありません: 現在の状態 = ${ws.readyState}`);
        return;
    }
    if (inputMessage.trim() !== "") {
        const message: Message = {
            content: inputMessage,
            senderID: "",
            recipientID: "",
            type: "self",
            username: "",
        };
        ws.send(JSON.stringify(message));
        console.log("メッセージを送信しました: ", message); // デバッグログ
        setInputMessage("");
    } else {
        console.log("入力されたメッセージが空です");
    }
};


  return (
    <Box sx={{ display: "flex", flexDirection: "column", width: "100%", m: 0 }}>
      <Box sx={{ mx: { md: 6 }, mb: 14 }}>
        <ChatBody data={messages} />
      </Box>
      <Box
        sx={{
          position: "fixed",
          bottom: 0,
          mt: 4,
          width: "100%",
          px: 2,
          py: 2,
          left: 0,
          bgcolor: "grey.100",
          mx: { md: 4 },
          borderRadius: "4px",
        }}
      >
        <Box
          sx={{
            display: "flex",
            flexDirection: { xs: "row", md: "row" },
            alignItems: "center",
          }}
        >
          <TextField
            multiline
            maxRows={4}
            placeholder="メッセージを入力"
            variant="outlined"
            sx={{
              width: "70%",
              mr: 2,
              borderRadius: "4px",
              borderColor: "primary.main",
              "& .MuiOutlinedInput-notchedOutline": {
                border: "1px solid",
                borderColor: "primary.main",
              },
            }}
            InputProps={{
              disableUnderline: true,
              style: { resize: "none", height: "40px" },
            }}
            value={inputMessage}
            onChange={(e) => setInputMessage(e.target.value)}
          />
          <Button
            variant="contained"
            color="primary"
            sx={{ px: 2, borderRadius: "4px" }}
            onClick={sendMessage}
          >
            送信
          </Button>
        </Box>
      </Box>
    </Box>
  );
};

export default ChatPage;
// "use client";
// import React, { useState, useRef, useContext } from "react";
// import { Box, TextField, Button } from "@mui/material";
// import { useRouter } from "next/navigation";
// import ChatBody from "../components/ChatBody";

// export type Message = {
//   content: string;
//   client_id: string;
//   username: string;
//   room_id: string;
//   type: "recv" | "self";
// };

// const index = () => {
//   const [messages, setMessage] = useState<Array<Message>>([
//     {
//       content: "よろしくお願いします",
//       client_id: "1",
//       username: "user1",
//       room_id: "1",
//       type: "self",
//     },
//     {
//       content: "hello",
//       client_id: "1",
//       username: "user1",
//       room_id: "1",
//       type: "self",
//     },
//     {
//       content: "hello",
//       client_id: "2",
//       username: "user2",
//       room_id: "1",
//       type: "recv",
//     },
//   ]);

//   return (
//     <Box sx={{ display: "flex", flexDirection: "column", width: "100%", m: 0 }}>
//       <Box sx={{ mx: { md: 6 }, mb: 14 }}>
//         <ChatBody data={messages} />
//       </Box>
//       <Box
//         sx={{
//           position: "fixed",
//           bottom: 0,
//           mt: 4,
//           width: "100%",
//           px: 2,
//           py: 2,
//           left: 0,
//           bgcolor: "grey.100",
//           mx: { md: 4 },
//           borderRadius: "4px",
//         }}
//       >
//         <Box
//           sx={{
//             display: "flex",
//             flexDirection: { xs: "row", md: "row" },
//             alignItems: "center",
//           }}
//         >
//           <TextField
//             multiline
//             maxRows={4}
//             placeholder="メッセージを入力"
//             variant="outlined"
//             sx={{
//               width: "70%",
//               mr: 2,
//               borderRadius: "4px",
//               borderColor: "primary.main",
//               "& .MuiOutlinedInput-notchedOutline": {
//                 border: "1px solid",
//                 borderColor: "primary.main", // Border color
//               },
//             }}
//             InputProps={{
//               disableUnderline: true, // Underlineを無効化
//               style: { resize: "none", height: "40px" }, // テキストエリアの高さを調整
//             }}
//           />
//           <Button
//             variant="contained"
//             color="primary"
//             sx={{ px: 2, borderRadius: "4px" }}
//           >
//             送信
//           </Button>
//         </Box>
//       </Box>
//     </Box>
//   );
// };

// export default index;
