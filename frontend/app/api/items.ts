// api/items.ts //なんのため？？？
import axios from "axios";
import { Item } from "../types/Items";

const API_URL = "http://localhost:8080";

export const fetchItems = async (): Promise<Item[]> => {
  try {
    const token = localStorage.getItem("token");
    const response = await axios.get<Item[]>(`${API_URL}/items`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    return response.data;
  } catch (error) {
    console.error("Error fetching items:", error);
    return [];
  }
};
