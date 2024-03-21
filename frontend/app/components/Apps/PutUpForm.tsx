import { useState } from "react";
import ItemMyList from "./ItemList/ItemMyList";
import { InputForm } from "./Listing/InputForm";

function App() {
  // reload ItemList after Listing complete
  const [reload, setReload] = useState(true);
  return (
    <div>
      <div>
        <InputForm onListingCompleted={() => setReload(true)} />
      </div>
      <div>
        <ItemMyList reload={reload} onLoadCompleted={() => setReload(false)} />
      </div>
    </div>
  );
}

export default App;
