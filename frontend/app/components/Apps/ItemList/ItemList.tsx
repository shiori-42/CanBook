// import React, { useEffect, useState } from "react";

// interface Item {
//   id: number;
//   image_name: string;
//   text_name: string;
//   class_name: string;
//   price: number;
//   sell_type: number;
// }

// const server = process.env.REACT_APP_API_URL || "http://127.0.0.1:9000";
// const placeholderImage = process.env.PUBLIC_URL + "/logo192.png";

// interface Prop {
//   reload?: boolean;
//   onLoadCompleted?: () => void;
// }

// export const ItemList: React.FC<Prop> = (props) => {
//   const { reload = true, onLoadCompleted } = props;
//   const [items, setItems] = useState<Item[]>([]);
//   const fetchItems = () => {
//     fetch(server.concat("/items"), {
//       method: "GET",
//       mode: "cors",
//       headers: {
//         "Content-Type": "application/json",
//         Accept: "application/json",
//       },
//     })
//       .then((response) => response.json())
//       .then((data) => {
//         console.log("GET success:", data);
//         setItems(data.items);
//         onLoadCompleted && onLoadCompleted();
//       })
//       .catch((error) => {
//         console.error("GET error:", error);
//       });
//   };

//   useEffect(() => {
//     if (reload) {
//       fetchItems();
//     }
//   }, [reload]);

//   return (
//     <div>
//       {items.map((item) => {
//         return (
//           <div key={item.id} className="ItemList">
//             {/* //Replace the placeholder image with the item image */}
//             <img src={placeholderImage} />
//             <p>
//               <span>Name: {item.text_name}</span>
//               <br />
//               <span>Category: {item.class_name}</span>
//             </p>
//           </div>
//         );
//       })}
//     </div>
//   );
// };
