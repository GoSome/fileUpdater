import PropTypes from "prop-types";
import { Box, Button, ListItem } from "@mui/material";

//redux
import { setName, setFilePath, increment } from "../updaters/updaterSlice";
import { useDispatch, useSelector } from "react-redux";

export const NavItem = props => {
  const { href, icon, title, filePath, ...others } = props;
  const name = useSelector(state => state.updater.name);
  const active = name ? title === name : false;
  const dispatch = useDispatch();

  return (
    <ListItem
      disableGutters
      sx={{
        display: "flex",
        mb: 0.5,
        py: 0,
        px: 2,
      }}
      {...others}
    >
      <Button
        component="a"
        startIcon={icon}
        disableRipple
        onClick={() => {
          dispatch(setName(title));
          dispatch(setFilePath(filePath));
          dispatch(increment());
        }}
        sx={{
          backgroundColor: active && "rgba(255,255,255, 0.08)",
          borderRadius: 1,
          color: active ? "secondary.main" : "neutral.300",
          fontWeight: active && "fontWeightBold",
          justifyContent: "flex-start",
          fontSize: 20,
          px: 8,
          textAlign: "left",
          textTransform: "none",
          width: "100%",
          "& .MuiButton-startIcon": {
            color: active ? "secondary.main" : "neutral.400",
          },
          "&:hover": {
            backgroundColor: "rgba(255,255,255, 0.08)",
          },
        }}
      >
        <Box sx={{ flexGrow: 1 }}>{title}</Box>
      </Button>
    </ListItem>
  );
};

NavItem.propTypes = {
  href: PropTypes.string,
  icon: PropTypes.node,
  title: PropTypes.string,
};
