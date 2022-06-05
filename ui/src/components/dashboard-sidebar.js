import { useEffect, useState } from "react";
// import { useRouter } from "next/router";
import PropTypes from "prop-types";
import { Box, Divider, Drawer, Typography, useMediaQuery } from "@mui/material";
import { NavItem } from "./nav-item";
import ArticleIcon from "@mui/icons-material/Article";
import { useSelector } from "react-redux";

const axios = require("axios");

const GetAllUpdatersURl = "http://192.168.2.5:8090/api/updaters";

export const DashboardSidebar = props => {
  const { open, onClose } = props;
  // const router = useRouter();
  const lgUp = useMediaQuery(theme => theme.breakpoints.up("lg"), {
    defaultMatches: true,
    noSsr: false,
  });
  const [updaters, setUpdaters] = useState([]);
  const value = useSelector(state => state.updater.value);
  const reFlash = useSelector(state => state.updater.reFlash);

  useEffect(() => {
    axios.get(GetAllUpdatersURl).then(function (response) {
      setUpdaters(response.data);
    });
  }, [reFlash]);

  useEffect(
    () => {
      if (!lgUp) {
        onClose();
      }
    },
    // eslint-disable-next-line react-hooks/exhaustive-deps
    [value]
  );

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
