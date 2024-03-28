"use client";
import * as React from "react";
import Typography from "@mui/material/Typography";
import { Box, Tabs, Tab, Stack, Button, IconButton } from "@mui/material";
// import ItemKeep from "./ItemKeep";
// import ItemAll from "./ItemAll";
import ItemMyList from "./Listing/ItemList/ItemMyList";
import ItemAll from "./ItemAll";
import SearchBox from "./SearchButton";
import SearchButton from "./SearchButton";
import { AccountCircle } from "@mui/icons-material";
import EditIcon from "@mui/icons-material/Edit";
import DeleteIcon from "@mui/icons-material/Delete";
import CollegeSearchButton from "./CollegeSearchButton";

interface TabPanelProps {
  children?: React.ReactNode;
  index: number;
  value: number;
}

function CustomTabPanel(props: TabPanelProps) {
  const { children, value, index, ...other } = props;

  return (
    <div
      role="tabpanel"
      hidden={value !== index}
      id={`simple-tabpanel-${index}`}
      aria-labelledby={`simple-tab-${index}`}
      {...other}
    >
      {value === index && <Box sx={{ pt: 1 }}>{children}</Box>}
    </div>
  );
}

function a11yProps(index: number) {
  return {
    id: `simple-tab-${index}`,
    "aria-controls": `simple-tabpanel-${index}`,
  };
}

export default function BasicTabs() {
  const [value, setValue] = React.useState(0);

  const handleChange = (event: React.SyntheticEvent, newValue: number) => {
    setValue(newValue);
  };

  return (
    <Box sx={{ width: "100%" }}>
      <Box sx={{ borderBottom: 1, borderColor: "divider" }}>
        <Tabs
          value={value}
          onChange={handleChange}
          aria-label="basic tabs example"
          centered
          sx={{
            mt: 2,
            "& .MuiTabs-indicator": {
              backgroundColor: "#009C88",
            },
          }}
        >
          <Tab
            label="すべて"
            {...a11yProps(0)}
            sx={{ "&.Mui-selected": { color: "#009C88" } }}
          />
          <Tab
            label="マイリスト"
            {...a11yProps(1)}
            sx={{ "&.Mui-selected": { color: "#009C88" } }}
          />
          <Tab
            label="保存済み"
            {...a11yProps(2)}
            sx={{ "&.Mui-selected": { color: "#009C88" } }}
          />
        </Tabs>
      </Box>
      <CustomTabPanel value={value} index={0}>
        <Stack direction="row" spacing={1} justifyContent="end" mt={1}>
          <SearchButton />
          <CollegeSearchButton />
        </Stack>
        <Typography fontSize={15} fontWeight={"bold"} mt={2}>
          すべての商品
        </Typography>
        <ItemAll />
      </CustomTabPanel>

      <CustomTabPanel value={value} index={1}>
        <Stack direction="row" spacing={2} justifyContent="end">
          <IconButton
            // aria-label="account of current user"
            // aria-controls="menu-appbar"
            // aria-haspopup="true"
            sx={{ color: "#F47381" }}
          >
            <EditIcon />
          </IconButton>
          <IconButton
            aria-label="account of current user"
            aria-controls="menu-appbar"
            aria-haspopup="true"
            sx={{ color: "#F47381" }}
          >
            <DeleteIcon />
          </IconButton>
        </Stack>
        <Typography fontSize={15} fontWeight={"bold"} mt={1.5}>
          出品した商品
        </Typography>
        <ItemMyList />
      </CustomTabPanel>

      <CustomTabPanel value={value} index={2}>
        <Typography fontSize={15} fontWeight={"bold"}>
          保存した商品
        </Typography>
        {/* <ItemKeep /> */}
      </CustomTabPanel>
    </Box>
  );
}
