//バックエンドからもらったトークン(token)をクライアント側のlocalStorageに保存している
//出品するときなどはそのtokenをバックエンド側に送って(headers)その認証ができたら出品ができる、という流れ
//現状どこのコンポーネントでも以下のコードを書いているが、このファイルを他のコンポーネントでも使い回すようにリファクタする
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
