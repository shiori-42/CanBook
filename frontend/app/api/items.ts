// api/items.ts //なんのため？？？
import axios from "axios";
import { item } from "../types/items";
const API_URL = "http://localhost:8080";

export const fetchItems = async (): Promise<item[]> => {
  try {
    const token = localStorage.getItem("token");
    const response = await axios.get<item[]>(`${API_URL}/items`, {
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
