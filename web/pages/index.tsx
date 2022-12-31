import Chart from '#/home/chart';
import Layout from '#/layout';
import { OverviewData, Overview } from '@/system/overview';
import { Container, Paper } from '@mui/material';
import Grid from '@mui/material/Unstable_Grid2';
import { useEffect, useState } from 'react';

export default function Home() {
  const [overviewData, setOverviewData] = useState<(OverviewData | void)[]>([]);

  useEffect(() => {
    Overview()?.then((data) => {
      if (data && !('error' in data)) setOverviewData(data);
    });

    const loop = setInterval(async () => {
      // TODO use limit param
      Overview()?.then((data) => {
        if (data && !('error' in data)) {
          setOverviewData([...data]);
        }
      });
    }, 10 * 1e3);
    return () => clearInterval(loop);
  }, []);

  const getValue = () => {
    const strTime = (date: Date) => {
      return [
        date.getHours().toString().padStart(2, '0'),
        date.getMinutes().toString().padStart(2, '0'),
        date.getSeconds().toString().padStart(2, '0'),
      ].join(':');
    };

    return overviewData
      .slice(0, 60)
      .reverse()
      .map((d) => {
        return (
          d && {
            time: strTime(new Date(d.time)),
            cpu_usage: d.cpu_usage,
            mem_usage: d.mem_usage,
          }
        );
      })
      .fill(void 0, 60);
  };

  return (
    <Layout>
      <Container maxWidth="lg">
        <Grid container spacing={2}>
          <Grid xs={6}>
            <Paper
              sx={{
                p: 2,
                display: 'flex',
                flexDirection: 'column',
                height: 240,
                color: 'white',
              }}
            >
              <Chart
                title="CPU 使用率"
                type={'cpu_usage'}
                data={getValue()}
                color="#56b0f5"
              />
            </Paper>
          </Grid>
          <Grid xs={6}>
            <Paper
              sx={{
                p: 2,
                display: 'flex',
                flexDirection: 'column',
                height: 240,
              }}
            >
              <Chart
                title="內存使用率"
                type={'mem_usage'}
                data={getValue()}
                color="#8884d8"
              />
            </Paper>
          </Grid>
        </Grid>
      </Container>
    </Layout>
  );
}
