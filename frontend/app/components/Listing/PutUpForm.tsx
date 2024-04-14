import { useState } from "react";
import InputForm from "./Form/InputForm";

function PutUpform() {
  //form/page.tsxで記述
  const [reload, setReload] = useState(true);
  return (
    <InputForm onListingCompleted={() => setReload(true)} />

    // {/* <ItemMyList reload={reload} onLoadCompleted={() => setReload(false)} /> */}
  );
}

export default PutUpform;
