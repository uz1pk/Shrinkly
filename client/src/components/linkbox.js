import { useState } from "react";

export default function LinkBox(props) {
  const linkToCopy = props.value;
  const [linkState, setlinkState] = useState("Copy");


  return (
    <div className="link-wrapper">
      <div>
        <p className="full-link">{`http://127.0.0.1:3001/shrinkly/${props.value}`}</p>
        <p className="short-link" id="short-link">
          {linkToCopy}
        </p>
      </div>
      <button
        className={linkState === "Copy" ? "button" : "button active"}
        onClick={() => {
          navigator.clipboard.writeText(`http://127.0.0.1:3001/r/${props.value}`);
          setlinkState("Copied!");
          }}
      >
        {linkState}
      </button>
    </div>
  );
}
