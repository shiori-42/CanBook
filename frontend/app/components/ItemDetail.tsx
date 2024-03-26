// "use client";
// import { Items } from "@/app/data/data";
// import NotFound from "@/app/not-found";
// import { Box, Stack, Typography, styled } from "@mui/material";
// import React, { Fragment } from "react";

// const ItemDetail = async ({ params }: { params: { id: string } }) => {
//   const item = Items.find((itemdata) => itemdata.id === parseInt(params.id));

//   if (!item) {
//     return <NotFound />;
//   }

//   // const backgroundImageStyle = {
//   //   backgroundImage: `url(${item.imagepath})`,
//   //   backgroundSize: "cover",
//   //   backgroundPosition: "center",
//   //   width: "100%",
//   //   height: 500,
//   // };

//   const StyledBox = styled(Box)(({ theme }) => ({
//     position: "relative",
//     overflow: "hidden",
//     backgroundSize: "cover",
//     backgroundPosition: "center",
//     backgroundImage: `url(${item.image_name})`,
//     // textAlign: "center",
//     width: "100%",
//     height: 480,
//     zIndex: 0,
//     [theme.breakpoints.down("sm")]: {
//       height: 400,
//     },
//   }));

//   return (
//     <>
//       <Box
//         width={"100%"}
//         height={"100%"}
//         borderBottom={1}
//         borderColor={"divider"}
//         pb={3}
//         // sx={backgroundImageStyle}
//       >
//         <StyledBox />
//         {/* <Typography fontSize={20} fontWeight={"bold"}>
//           {item.text_name}
//         </Typography>
//         <Typography fontSize={20}>￥{item.price}</Typography>
//         <Typography fontSize={12}>{item.sell_type}</Typography> */}
//         <Typography
//           fontSize={{ xs: 18, sm: 20 }}
//           height={40}
//           px={0.5}
//           mt={1}
//           fontWeight={"bold"}
//         >
//           {item.text_name}
//         </Typography>
//         <Stack direction={"row"} mt={1}>
//           <Typography
//             fontSize={{ xs: 9, sm: 11 }}
//             mx={0.5}
//             mt={{ xs: 0.7, sm: 0.5 }}
//             fontWeight={"bold"}
//             textAlign={"center"}
//             justifyContent={"center"}
//             width={{ xs: 49, sm: 100 }}
//             height={{ xs: 15, sm: 18 }}
//             border={1.9}
//             borderRadius={"20px"}
//             color={item.sell_type === "レンタル" ? "#009C88" : "#1573FF"}
//             style={{ display: "flex", alignItems: "center" }}
//           >
//             {item.sell_type}
//           </Typography>
//           <Typography
//             fontSize={{ xs: 20, sm: 25 }}
//             width={100}
//             alignItems={"center"}
//             justifyContent={"center"}
//             color={item.sell_type === "レンタル" ? "#009C88" : "#1573FF"}
//           >
//             ￥{item.price}
//           </Typography>
//         </Stack>
//       </Box>
//       {/* <Typography fontSize={15} fontWeight={"bold"} mt={1}>
//         商品の詳細
//       </Typography>
//       <ItemInfo /> */}
//     </>
//   );
// };

// export default ItemDetail;
