using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;
using System.Diagnostics;
namespace Promt
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();
        }

        private void label1_Click(object sender, EventArgs e)
        {

        }

        private void Form1_Load(object sender, EventArgs e)
        {

        }

        private void button1_Click(object sender, EventArgs e)
        {
            //Process.Start(@"C:\Windows\System32\cmd.exe");
            Application.Exit();
        }

        int count = 100;
        private void timerCall(object sender, EventArgs e)
        {
            if (count <= 0)
            {
                timer1.Enabled = false;
                msgBox();
            }
            else
            {
                countDown.Text = parseTime(count--);
            }
        }
        string parseTime(int secs)
        {
            string s;
            int min = 0;
            int sec = 0;
            int hrs = 0;
            int dys = 0;
            sec = secs;

            if (sec > 60)
            {
                min = sec / 60;
                sec -= min * 60;
            }
            if (min > 60)
            {
                hrs = min / 60;
                min -= hrs * 60;
            }
            if (hrs>24)
            {
                dys = hrs / 24;
                hrs -= dys * 24;
            }
            //Format
            s = string.Format("{0} Days, {1} Hours, {2} Minutes, {3} Seconds.", dys, hrs, min, sec);

            return s;
        }
        private void msgBox()
        {
            string message = "You did not enter a server name. Cancel this operation?";
            string caption = "Error Detected in Input";
            MessageBoxButtons buttons = MessageBoxButtons.YesNo;
            DialogResult result;

            // Displays the MessageBox.
     
            result = MessageBox.Show(message, caption, buttons);
      
            if (result == System.Windows.Forms.DialogResult.Yes)
                    this.Close();

                }

        private void countDown_Click(object sender, EventArgs e)
        {

        }

        private void Info_Click(object sender, EventArgs e)
        {

        }

        private void label3_Click(object sender, EventArgs e)
        {

        }
    }

}

