"use client";
import { Box, Typography, styled } from "@mui/material";

const Profile = () => {
  const StyledBox = styled(Box)(({ theme }) => ({
    overflow: "hidden",
    backgroundSize: "cover",
    backgroundPosition: "center",
    backgroundImage: `url(/Dots-to-Code.png)`,
    width: 200,
    height: 200,
    borderRadius: "50%",
    [theme.breakpoints.down("sm")]: {
      width: 100,
      height: 100,
    },
  }));

  return (
    <>
      <Box
        width={"100%"}
        height={"100%"}
        borderBottom={1}
        borderColor={"divider"}
        mt={3}
        pb={3}
        display="flex"
        flexDirection="column"
        alignItems="center"
      >
        <StyledBox />
        <Typography fontSize={20} fontWeight={"bold"} mt={1.5}>
          ああああああ
        </Typography>
      </Box>

      <Typography fontSize={20} fontWeight={"bold"}>
        あああ
      </Typography>
      <Typography fontSize={20}>あああ</Typography>
      <Typography fontSize={12}>あああ</Typography>
    </>
  );
};

export default Profile;
