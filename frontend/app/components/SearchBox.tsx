"use client";
import { alpha, styled } from "@mui/material/styles";
import InputBase from "@mui/material/InputBase";
import SearchIcon from "@mui/icons-material/Search";
import { useState } from "react";

const Search = styled("div")(({ theme }) => ({
  position: "relative",
  marginRight: theme.spacing(2),
  marginLeft: 0,
  marginTop: 18,
  width: "100%",

  border: `1px solid #ABABAB`,
}));

const SearchIconWrapper = styled("div")(({ theme }) => ({
  padding: theme.spacing(0, 2),
  height: "100%",
  position: "absolute",
  pointerEvents: "none",
  display: "flex",
  alignItems: "center",
  justifyContent: "center",
}));

const StyledInputBase = styled(InputBase)(({ theme }) => ({
  color: "inherit",
  "& .MuiInputBase-input": {
    padding: theme.spacing(1, 1, 1, 0),
    paddingLeft: `calc(1em + ${theme.spacing(4)})`,
    transition: theme.transitions.create("width"),
    width: "100%",
  },
}));

// デフォルトの値として空の関数を設定しておく
export default function SearchBox({ onSearch = () => {} }) {
  const [searchQuery, setSearchQuery] = useState("");

  const handleInputChange = (event: any) => {
    const { value } = event.target;
    setSearchQuery(value);
    // onSearch プロパティが関数であることを確認してから呼び出す
    if (typeof onSearch === "function") {
      onSearch();
    }
  };

  return (
    <Search>
      <SearchIconWrapper>
        <SearchIcon />
      </SearchIconWrapper>
      <StyledInputBase
        placeholder="何をお探しですか"
        inputProps={{ "aria-label": "search" }}
        value={searchQuery}
        onChange={handleInputChange}
      />
    </Search>
  );
}
