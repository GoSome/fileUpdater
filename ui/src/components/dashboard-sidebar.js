import { useEffect, useState } from "react";
// import { useRouter } from "next/router";
import PropTypes from "prop-types";
import {
  Box,
  Button,
  Divider,
  Drawer,
  Typography,
  useMediaQuery,
} from "@mui/material";
import { NavItem } from "./nav-item";
import ArticleIcon from "@mui/icons-material/Article";
const axios = require("axios");

const GetAllUpdatersURl = "http://192.168.2.5:8090/api/updaters";
const UpdateFileURl = "http://192.168.2.5:8090/api/content";

export const DashboardSidebar = props => {
  const { open, onClose } = props;
  // const router = useRouter();
  const lgUp = useMediaQuery(theme => theme.breakpoints.up("lg"), {
    defaultMatches: true,
    noSsr: false,
  });
  const [updaters, setUpdaters] = useState([]);

  useEffect(() => {
    axios.get(GetAllUpdatersURl).then(function (response) {
      console.log("get all updaters: ", response.data);
      setUpdaters(response.data);
    });
  }, []);

  // useEffect(
  //   () => {
  //     if (!router.isReady) {
  //       return;
  //     }

  //     if (open) {
  //       onClose?.();
  //     }
  //   },
  //   // eslint-disable-next-line react-hooks/exhaustive-deps
  //   [router.asPath]
  // );

  const content = (
    <>
      <Box
        sx={{
          display: "flex",
          flexDirection: "column",
          height: "100%",
        }}
      >
        <Box sx={{ my: 2 }}>
          <Typography variant="h4" sx={{ textAlign: "center" }}>
            FileUpdater
          </Typography>
        </Box>
        <Divider
          sx={{
            borderColor: "#111828",
            my: 0,
            mb: 2,
          }}
        />
        <Box sx={{ flexGrow: 1 }}>
          {updaters.map(item => (
            <NavItem
              key={item.name}
              title={item.name}
              icon={<ArticleIcon fontSize="small" />}
              filePath={item.path}
            />
          ))}
        </Box>
      </Box>
    </>
  );

  if (lgUp) {
    return (
      <Drawer
        anchor="left"
        open
        PaperProps={{
          sx: {
            backgroundColor: "neutral.900",
            color: "#FFFFFF",
            width: 280,
          },
        }}
        variant="permanent"
      >
        {content}
      </Drawer>
    );
  }

  return (
    <Drawer
      anchor="left"
      onClose={onClose}
      open={open}
      PaperProps={{
        sx: {
          backgroundColor: "neutral.900",
          color: "#FFFFFF",
          width: 280,
        },
      }}
      sx={{ zIndex: theme => theme.zIndex.appBar + 100 }}
      variant="temporary"
    >
      {content}
    </Drawer>
  );
};

DashboardSidebar.propTypes = {
  onClose: PropTypes.func,
  open: PropTypes.bool,
};
