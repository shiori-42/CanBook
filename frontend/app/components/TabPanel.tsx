import * as React from "react";
import Typography from "@mui/material/Typography";
import ItemGrid from "./ItemAll";
import { Box, Tabs, Tab } from "@mui/material";
import ItemDetail from "./ItemDetail";
import ItemKeep from "./ItemKeep";
import ItemMyList from "./ItemMyList";

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
            "& .MuiTabs-indicator": {
              backgroundColor: "#32C0C9",
            },
          }}
        >
          <Tab
            label="キャンパス"
            {...a11yProps(0)}
            sx={{ "&.Mui-selected": { color: "#32C0C9" } }}
          />
          <Tab
            label="マイリスト"
            {...a11yProps(1)}
            sx={{ "&.Mui-selected": { color: "#32C0C9" } }}
          />
          <Tab
            label="保存済み"
            {...a11yProps(2)}
            sx={{ "&.Mui-selected": { color: "#32C0C9" } }}
          />
        </Tabs>
      </Box>
      <CustomTabPanel value={value} index={0}>
        <Typography fontSize={15} fontWeight={"bold"}>
          キャンパス内の商品
        </Typography>
        <ItemGrid />
      </CustomTabPanel>
      <CustomTabPanel value={value} index={1}>
        <Typography fontSize={15} fontWeight={"bold"}>
          出品した商品
        </Typography>
        <ItemMyList />
      </CustomTabPanel>
      <CustomTabPanel value={value} index={2}>
        <Typography fontSize={15} fontWeight={"bold"}>
          保存した商品
        </Typography>
        <ItemKeep />
      </CustomTabPanel>
    </Box>
  );
}
