import React from "react";
import { useEffect } from "react";
import "@uiw/react-textarea-code-editor/dist.css";
//redux
import { useDispatch, useSelector } from "react-redux";
import { setContent } from "../../updaters/updaterSlice";
import CodeEditor from "@uiw/react-textarea-code-editor";

const axios = require("axios").default;

function MyCodeEditor(props) {
  const dispatch = useDispatch();
  const fileName = useSelector(state => state.updater.name);
  const language = useSelector(state => state.updater.language);
  const content = useSelector(state => state.updater.content);


  useEffect(() => {
    async function getContent() {
      try {
        const res = await axios.get("/api/content", {
          params: { name: fileName },
        });
        const { content } = await res.data;
        dispatch(setContent(content));
      } catch (error) {
        console.error(error);
      }
    }
    getContent();
  }, [fileName, dispatch]);

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
          fontSize: 16,
          backgroundColor: "#111827",
          color: "white",
          fontFamily:
            "ui-monospace,SFMono-Regular,SF Mono,Consolas,Liberation Mono,Menlo,monospace",
        }}
      />
    </div>
  );
}

export default MyCodeEditor;
