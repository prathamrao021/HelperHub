import * as React from 'react';
import { AppBar, Toolbar, Typography, Button, Box, Drawer, IconButton, TextField, Switch } from '@mui/material';
import MenuIcon from '@mui/icons-material/Menu';
import { Link } from 'react-router-dom'; 
import { useTheme } from '@mui/material/styles';

export default function Navbar() {
  const [drawerOpen, setDrawerOpen] = React.useState(false);
  const [darkMode, setDarkMode] = React.useState(false);
  const theme = useTheme();

  const toggleDrawer = (open: boolean) => {
    setDrawerOpen(open);
  };

  const handleThemeChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setDarkMode(event.target.checked);
  };

  return (
    <AppBar position="sticky" color="primary">
      <Toolbar>
        <Typography variant="h6" sx={{ flexGrow: 1 }}>
          HelperHub
        </Typography>

        {/* Search Bar */}
        <TextField
          variant="outlined"
          placeholder="Search..."
          size="small"
          sx={{
            marginRight: 2,
            backgroundColor: theme.palette.background.paper,
            borderRadius: 1,
            '& .MuiInputBase-root': {
              paddingLeft: 2,
            },
          }}
        />

        {/* Desktop Navigation */}
        <Box sx={{ display: { xs: 'none', sm: 'flex' } }}>
          <Button color="inherit" component={Link} to="/home">
            Home
          </Button>
          <Button color="inherit" component={Link} to="/about">
            About
          </Button>
          <Button color="inherit" component={Link} to="/contact">
            Contact
          </Button>
        </Box>

        {/* Toggle Dark/Light Mode */}
        <Switch checked={darkMode} onChange={handleThemeChange} />

        {/* Mobile Menu Icon */}
        <IconButton
          color="inherit"
          aria-label="menu"
          edge="end"
          onClick={() => toggleDrawer(true)}
          sx={{ display: { xs: 'block', sm: 'none' } }}
        >
          <MenuIcon />
        </IconButton>
      </Toolbar>

      {/* Drawer for Mobile */}
      <Drawer
        anchor="right"
        open={drawerOpen}
        onClose={() => toggleDrawer(false)}
      >
        <Box sx={{ width: 250 }} role="presentation" onClick={() => toggleDrawer(false)}>
          <Button color="inherit" fullWidth component={Link} to="/home">
            Home
          </Button>
          <Button color="inherit" fullWidth component={Link} to="/about">
            About
          </Button>
          <Button color="inherit" fullWidth component={Link} to="/contact">
            Contact
          </Button>
        </Box>
      </Drawer>
    </AppBar>
  );
}
