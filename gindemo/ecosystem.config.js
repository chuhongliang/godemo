const production = process.env.NODE_ENV === 'production';

module.exports = {
  apps : [{
    name: 'gindemo',
    port: 7001,
    script: './server.exe',
    instances: production ? 1 : 1,
    autorestart: true,
    watch: false,
    exec_mode: 'cluster',
    restart_delay: 3000,
    exp_backoff_restart_delay: 100,
    max_memory_restart: '4G',
    cron_restart: '',
    time: true,
    error_file: './logs/pm2_error.log',
    out_file: './logs/pm2_out.log',
    log_file: './logs/pm2_combined.log',
    env: {
      NODE_ENV: 'development',
    },
    env_production: {
      NODE_ENV: 'production',
    }
  }]
};
