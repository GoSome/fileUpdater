import {
  Box,
  Button,
  Card,
  CardHeader,
  Divider,
  InputLabel,
  NativeSelect,
  Snackbar,
  Alert,
} from "@mui/material";

import FormControl from "@mui/material/FormControl";
import MyCodeEditor from "./Editor";
import { useSelector, useDispatch } from "react-redux";
import { setLanguage, setReFlash } from "../../updaters/updaterSlice";
import { useState } from "react";

const axios = require("axios").default;

const EditorLanguages = ["plaintext", "ini", "yaml", "json"];

export const EditorMain = props => {
  const [alertContent, setAlertContent] = useState("saved");
  const [level, setLevel] = useState("success");
  const [open, setOpen] = useState(false);

  const RaiseAlert = () => {
    setOpen(true);
  };

  const handleClose = (event, reason) => {
    if (reason === "clickaway") {
      return;
    }
    setOpen(false);
  };

  const name = useSelector(state => state.updater.name);
  const filePath = useSelector(state => state.updater.filePath);

  const content = useSelector(state => state.updater.content);

  const dispatch = useDispatch();
  const handleSave = e => {
    e.preventDefault();

    const j = JSON.stringify({ name: name, content: content });
    axios
      .post("/api/content", j)
      .then(function (response) {
        if (response.status === 200) {
          setAlertContent("saved");
          setLevel("success");
          RaiseAlert();
          dispatch(setReFlash());
        } else {
          setAlertContent("save file failed! status:" + response.status);
          setLevel("error");
          RaiseAlert();
        }
      })
      .catch(function (error) {
        setAlertContent("error! " + error);
        setLevel("error");
        RaiseAlert();
      })
      .then(function () {
        // always executed
      });
  };

  return (
    <form {...props}>
      <Card>
        <CardHeader
          subheader={filePath}
          title={name}
          sx={{
            display: "flex",
            justifyContent: "center",
            fontSize: "3rem",
            p: 1,
          }}
        />
        <Divider />
        <Box
          sx={{
            display: "flex",
            justifyContent: "flex-left",
            p: 2,
          }}
        >
          <FormControl>
            <InputLabel variant="standard" htmlFor="uncontrolled-native">
              language
            </InputLabel>
            <NativeSelect
              onChange={evn => {
                dispatch(setLanguage(evn.target.value));
              }}
            >
              {EditorLanguages.map(lang => (
                <option key={lang} value={lang}>
                  {lang}{" "}
                </option>
              ))}
            </NativeSelect>
          </FormControl>
        </Box>
        <MyCodeEditor />
        <Divider />
        <Box
          sx={{
            display: "flex",
            justifyContent: "flex-left",
            p: 2,
          }}
        >
          <Button
            color="primary"
            variant="contained"
            onClick={handleSave}
            sx={{
              backgroundColor: "#0CB982",
            }}
          >
            Save
          </Button>
          <Snackbar
            open={open}
            autoHideDuration={2000}
            onClose={handleClose}
            message="Saved"
            anchorOrigin={{ vertical: "top", horizontal: "center" }}
          >
            <Alert severity={level}>{alertContent}</Alert>
          </Snackbar>
        </Box>
      </Card>
    </form>
  );
};
