import { useState } from "react";
import InputForm from "./Form/InputForm";
import ItemMyList from "./ItemList/ItemMyList";

function PutUpform() {
  //form/page.tsxで記述
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

export default PutUpform;
