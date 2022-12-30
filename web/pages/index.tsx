import Chart from '#/home/chart';
import Layout from '#/layout';
import { Container, Grid, Paper } from '@mui/material';

export default function Home() {
  return (
    <Layout>
      <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
        <Grid container spacing={3}>
          <Grid item xs={12} mb={8} lg={9}>
            <Paper
              sx={{
                p: 2,
                display: 'flex',
                flexDirection: 'column',
                height: 240,
              }}
            >
              <Chart />
            </Paper>
          </Grid>
        </Grid>
      </Container>
    </Layout>
  );
}
