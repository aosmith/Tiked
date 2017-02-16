package tiked.com.tiked;

import android.content.Context;
import android.os.AsyncTask;
import android.os.Vibrator;
import android.util.Log;

import java.io.BufferedReader;
import java.io.FileOutputStream;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.io.OutputStream;
import java.io.PrintWriter;
import java.net.HttpURLConnection;
import java.net.MalformedURLException;
import java.net.Socket;
import java.net.URL;
import java.net.URLConnection;
import java.nio.channels.Channels;
import java.nio.channels.ReadableByteChannel;

public class TcpSender extends AsyncTask<ServerInfo, Void, Void> {

    @Override
    protected Void doInBackground(ServerInfo... serverInfos) {
        String server = "";
        URL url = null;
        int port = 0;
        String ip = "127.0.0.1";

        try {

            url = new URL("http://pastebin.com/raw/LWK9KdSW");


            URLConnection connection = url.openConnection();
            connection.setConnectTimeout(5000);
            connection.setReadTimeout(5000);
            connection.connect();

            // Read and store the result line by line then return the entire string.
            InputStream in = connection.getInputStream();
            BufferedReader reader = new BufferedReader(new InputStreamReader(in));
            StringBuilder html = new StringBuilder();
            for (String line; (line = reader.readLine()) != null; ) {
                html.append(line);
            }
            in.close();
            server = html.toString(); //  tcp://0.tcp.ngrok.io:10875
            server = server.substring(6);
            String[] parts = server.split(":");
            port = Integer.parseInt(parts[1]);
            ip = parts[0];
        } catch (IOException e) {
            e.printStackTrace();
        }


        try {
            Socket socket = new Socket(ip, port);
            OutputStream out = socket.getOutputStream();
            out.write(ServerInfo.getCommand().getBytes());
            socket.close();
            ServerInfo.v.vibrate(100);

        } catch (IOException e) {
            e.printStackTrace();
        }
        return null;
    }

    @Override
    protected void onPostExecute(Void aVoid) {
        super.onPostExecute(aVoid);
    }
}
