import {
  Box,
  Grid,
  Typography,
  Button,
  Avatar,
  Paper,
  Stack,
  useTheme
} from '@mui/material';
import { Public } from '@mui/icons-material';

export const HeroSection = () => {
  const theme = useTheme();

  return (
    <Box
      sx={{
        height: '100%',
        position: 'relative',
        overflow: 'hidden',
        background: 'linear-gradient(45deg, #f8f9fa 0%, #e9ecef 100%)'
      }}
    >
      <Grid
        container
        spacing={6}
        sx={{
          maxWidth: 1218,
          margin: '0 auto',
          height: '100%',
          alignItems: { lg: 'flex-end' },
          flexDirection: { xs: 'column', lg: 'row' },
          justifyContent: { lg: 'space-between' }
        }}
      >
        {/* Left Content */}
        <Grid item xs={12} lg={6}>
          <Box
            sx={{
              display: 'flex',
              flexDirection: 'column',
              alignItems: { xs: 'center', md: 'flex-start' },
              textAlign: { xs: 'center', md: 'left' },
              position: 'relative',
              zIndex: 20
            }}
          >
            <Paper
              elevation={0}
              sx={{
                display: 'inline-flex',
                alignItems: 'center',
                gap: 2,
                p: 0.5,
                pr: 4,
                mb: 4,
                borderRadius: 28,
                bgcolor: 'background.paper',
                boxShadow: theme.shadows[1]
              }}
            >
              <Button
                variant="contained"
                sx={{
                  borderRadius: 28,
                  px: { xs: 2, md: 4 },
                  py: 1,
                  textTransform: 'none'
                }}
              >
                Join the Movement
              </Button>
              <Typography variant="body1" sx={{ maxWidth: 260 }}>
                <Public sx={{ verticalAlign: 'middle', mr: 1 }} />
                Connect with Causes that Inspire You at HelperHub
              </Typography>
            </Paper>

            <Typography
              variant="h1"
              sx={{
              fontWeight: 900,
              mb: 4,
              fontSize: {
                xs: '2.5rem',
                md: '3.5rem',
                lg: '4.5rem',
                xl: '5.5rem'
              },
              lineHeight: 1.2,
              color: 'text.primary'
              }}
            >
              Empower Change Through Volunteering
            </Typography>

            <Typography
              variant="h5"
              sx={{
                mb: 6,
                color: 'text.secondary',
                maxWidth: 600,
                lineHeight: 1.6
              }}
            >
              Find your passion and make a difference with our tailored volunteer matching platform.
            </Typography>

            <Button
              variant="contained"
              size="large"
              href="/register/volunteer"
              sx={{
                px: 6,
                py: 2,
                fontSize: '1.25rem',
                borderRadius: 2,
                textTransform: 'none'
              }}
            >
              Become a Volunteer
            </Button>
          </Box>
        </Grid>

        {/* Right Content */}
        <Grid item xs={12} lg={5}>
          <Paper
            elevation={8}
            sx={{
              p: 4,
              borderRadius: 6,
              border: '4px solid',
              borderColor: 'divider',
              background: 'rgba(255, 255, 255, 0.4)',
              backdropFilter: 'blur(20px)',
              position: 'relative',
              zIndex: 20
            }}
          >
            <Stack
              direction="row"
              spacing={-2}
              sx={{ mb: 4, justifyContent: 'center' }}
            >
              {[1, 2, 3, 4, 5].map((item) => (
              <Avatar
                key={item}
                src={`./src/assets/images/prof${item}.jpg`}
                sx={{
                width: 56,
                height: 56,
                border: '2px solid white'
                }}
              />
              ))}
            </Stack>

            <Typography
              variant="h3"
              sx={{
                textAlign: 'center',
                fontWeight: 600,
                '& span': {
                  color: 'primary.main',
                  fontSize: '2.5rem',
                  [theme.breakpoints.up('lg')]: {
                    fontSize: '3rem'
                  }
                }
              }}
            >
              <span>5,000+</span> Volunteers Ready to Help
            </Typography>
          </Paper>
        </Grid>
      </Grid>
    </Box>
  );
};