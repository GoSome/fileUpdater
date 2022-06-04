import React from "react";
import { useEffect, useState } from "react";
//redux
import { useDispatch, useSelector } from "react-redux";
import { setContent } from "../../updaters/updaterSlice";
import CodeEditor from "@uiw/react-textarea-code-editor";

const axios = require("axios").default;

function MyCodeEditor(props) {
  const dispatch = useDispatch();
  const name = useSelector(state => state.updater.name);
  const language = useSelector(state => state.updater.language);
  const content = useSelector(state => state.updater.content);

  useEffect(() => {
    axios
      .get("http://192.168.2.5:8090/api/content", {
        params: {
          name: name,
        },
      })
      .then(function (response) {
        console.log(name, response.data.content);
        dispatch(setContent(response.data.content));
      })
      .catch(function (error) {
        console.log(error);
      })
      .then(function () {
        // always executed
      });
  }, [name]);

  return (
    <div>
      <CodeEditor
        value={content}
        language={language}
        placeholder="please select file from left!"
        onChange={evn => dispatch(setContent(evn.target.value))}
        autoFocus
        padding={15}
        disabled={props.disabled}
        style={{
          fontSize: 14,
          backgroundColor: "#f5f5f5",
          fontFamily:
            "ui-monospace,SFMono-Regular,SF Mono,Consolas,Liberation Mono,Menlo,monospace",
        }}
      />
    </div>
  );
}

export default MyCodeEditor;
