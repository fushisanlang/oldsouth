import matplotlib.pyplot as plt
from matplotlib.backends.backend_tkagg import FigureCanvasTkAgg
import tkinter as tk
from tkinter import ttk
import requests
import datetime
import matplotlib.dates as mdates

class PlotApp:
    def __init__(self, root):
        self.root = root
        self.root.title("Real-Time Data Plot")

        # List of metacode values to cycle through
        self.metacodes = ["PBG", "W", "FBG"]
        self.current_metacode_index = 0

        # Create a frame for the plot
        self.plot_frame = ttk.Frame(self.root)
        self.plot_frame.pack(fill=tk.BOTH, expand=True)

        # Generate the initial plot
        self.fig, self.ax = plt.subplots()
        self.canvas = FigureCanvasTkAgg(self.fig, master=self.plot_frame)
        self.canvas.get_tk_widget().pack(fill=tk.BOTH, expand=True)

        # Start the auto update
        self.update_plot()

    def fetch_data(self, metacode):
        url = "http://10.1.75.228:8080/metas"
        # url= "http://10.1.75.224:5000/oldsouth/line"
        headers = {
            "token": "2Ux~BOIBCDU+6!hgDCp8",  # Replace with your actual token
            "metacode": metacode,  # Current metacode
            # "starttime":str(int((datetime.datetime.now() - datetime.timedelta(days=7)).timestamp())),
            # "stoptime":str(int(datetime.datetime.now().timestamp()))
        }
        try:
            response = requests.get(url, headers=headers)
            response.raise_for_status()
            data = response.json()
            print(response.status_code)
            return data["TimeList"], data["DataList"], data["Title"]
        except requests.RequestException as e:
            print(response.status_code)
            print(f"Error fetching data: {e}")
            return [], [], ""

    def plot_data(self, time_list, data_list, title):
        self.ax.clear()
        
        # Convert the time strings to datetime objects for proper plotting
        time_list = [datetime.datetime.strptime(t, "%Y%m%d-%H") for t in time_list]
        
        self.ax.plot(time_list, data_list, marker='o')
        self.ax.set_title(title)
        self.ax.set_xlabel('Time')
        self.ax.set_ylabel('Value')
        
        # Format x-axis for 7-day range
        self.ax.set_xlim([datetime.datetime.now() - datetime.timedelta(days=7), datetime.datetime.now()])
        self.ax.xaxis.set_major_locator(mdates.DayLocator())  # Major ticks every day
        self.ax.xaxis.set_major_formatter(mdates.DateFormatter('%Y-%m-%d'))
        self.ax.xaxis.set_minor_locator(mdates.HourLocator(interval=6))  # Minor ticks every 6 hours

        self.fig.autofmt_xdate()  # Auto-format the x-axis to show dates nicely
        
        self.canvas.draw()

    def update_plot(self):
        # Get the current metacode
        current_metacode = self.metacodes[self.current_metacode_index]
        #print(current_metacode)
        time_list, data_list, title = self.fetch_data(current_metacode)
        print(data_list)
        # print("Current Metacode:", current_metacode)  # Debugging print statement
        if time_list and data_list:
            self.plot_data(time_list, data_list, title)
        
        # Update to the next metacode in the list
        self.current_metacode_index = (self.current_metacode_index + 1) % len(self.metacodes)
        print("Next Metacode Index:", self.current_metacode_index)  # Debugging print statement
        
        # Schedule the next update
        self.root.after(1000, self.update_plot)  # Update every 10 seconds

# Create the main window
root = tk.Tk()
app = PlotApp(root)
root.mainloop()

