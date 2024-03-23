// import React, { useState } from "react";
// import Grid from "@mui/material/Grid";
// import { Box, Card, CardMedia, Typography } from "@mui/material";
// import Link from "next/link";
// import { items } from "../data/data";
// import SearchBox from "./SearchBox";

// const ItemAll = () => {
//   const [filteredItems, setFilteredItems] = useState(items);

//   // const handleSearch = (searchQuery: string) => {
//   //   const filtered = items.filter((item) =>
//   //     item.name.toLowerCase().includes(searchQuery.toLowerCase())
//   //   );
//   //   setFilteredItems(filtered);
//   // };　　　　　　　　　　　　　　　　　　　　　　　　　　　　　　　　　　//絞り込みしてる

//   return (
//     <>
//       <SearchBox onSearch={undefined} />
//       <Grid container spacing={1} py={2}>
//         {filteredItems.map((item, index) => (
//           <Grid key={index} item xs={4}>
//             <Link href={`/item/${item.id}`} style={{ textDecoration: "none" }}>
//               <Card
//                 sx={{
//                   height: { xs: 200, sm: 270 },
//                   transition: "0.3s",
//                   "&:hover": {
//                     bgcolor: "#ffffff",
//                     opacity: 0.5,
//                     transition: "0.3s",
//                   },
//                 }}
//               >
//                 <Box bgcolor={"#ededed"}>
//                   <CardMedia
//                     // sx={{ objectFit: "contain" }}
//                     component="img"
//                     sx={{ height: { xs: 110, sm: 180 } }}
//                     image={item.imagepath}
//                   />
//                 </Box>
//                 <Typography fontSize={13} height={40}>
//                   {item.name}
//                 </Typography>
//                 <Typography>￥{item.price}</Typography>
//                 <Typography fontSize={9}>{item.sell}</Typography>
//               </Card>
//             </Link>
//           </Grid>
//         ))}
//       </Grid>
//     </>
//   );
// };

// export default ItemAll;
