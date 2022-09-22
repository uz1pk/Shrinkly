import { useState } from "react";

import LinkBox from "./linkbox";
import { ReactComponent as MainImage } from "../images/main-screen-image.svg";

export default function Home() {
  //link shortening
  const [linkToShorten, setLinkToShorten] = useState();
  const [shortenedLinks, setShortenedLinks] = useState([]);

  // useEffect(() => {
  //   sessionStorage.setItem("ssLinks", JSON.stringify(shortenedLinks));
  // }, [shortenedLinks]);

  // function getSessionStorageOrDefault(key, defaultValue) {
  //   const stored = sessionStorage.getItem(key);
  //   if (!stored) {
  //     return defaultValue;
  //   }
  //   return JSON.parse(stored);
  // }

  function handleInput(el) {
    setLinkToShorten(el.target.value);
  }

  function handleSubmit() {
    if (linkToShorten === "" || linkToShorten === undefined) {
      alert("Invalid URL");
      return;
    }

    const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ shrinkly: "", redirect: linkToShorten, random: true })
    };
    
    fetch("http://127.0.0.1:3001/shrinkly", requestOptions)
    .then(res => res.json())
    .then(data => {
      console.log(data);
      setShortenedLinks((prevLinks) => [...prevLinks, data.shrinkly]);
    });
    console.log(shortenedLinks);
    setLinkToShorten("");
  }

  return (
    <>

        <div className="image-section">
          <MainImage className="image" />
          <div className="heading-section">
            <h1 className="fs-primary-heading">
              Shrinkly
            </h1>
            <p>
              The best new URL shortening service.
            </p>
          </div>
        </div>

        <div
          className="submit-link-section">
          <input
            type="text"
            name="link-input-field"
            className="link-input-field"
            placeholder="Paste link here..."
            value={linkToShorten || ""}
            onChange={handleInput}
          />
          <button
            className="button submit-btn"
            onClick={handleSubmit}
          >
            Shrinklify!
          </button>
        </div>

        <div className="generated-links">
          {shortenedLinks.map((link) => (
            <LinkBox value={link} />
          ))}
        </div>


    </>
  );
}