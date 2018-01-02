from crontab import CronTab

# File name for cron
my_cron = CronTab(tabfile='my_cron.tab')

# Add cron command and time span
job  = my_cron.new(command='python /usr/src/app/hello.py')
job.minute.every(1)

# Write cron jobs to cron tab file
my_cron.write()

# Run the scheduler
for result in my_cron.run_scheduler():
    print ("This was printed to stdout by the process.")
