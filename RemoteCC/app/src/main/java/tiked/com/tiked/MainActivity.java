package tiked.com.tiked;

import android.content.Context;
import android.content.Intent;
import android.os.AsyncTask;
import android.os.Bundle;
import android.os.Vibrator;
import android.support.v7.app.AppCompatActivity;
import android.util.Log;
import android.view.View;
import android.widget.EditText;
import android.widget.Switch;

import com.google.android.gms.wearable.MessageEvent;

import static tiked.com.tiked.ServerInfo.v;


public class MainActivity extends AppCompatActivity {

    EditText argt, cmdt, trgt;
    Switch sw;
    Boolean swState = false;



    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        argt = (EditText) findViewById(R.id.args);
        cmdt = (EditText) findViewById(R.id.cmdt);
        trgt = (EditText) findViewById(R.id.trgt);
        sw = (Switch) findViewById(R.id.switch1);

        ServerInfo.setV((Vibrator) this.getSystemService(Context.VIBRATOR_SERVICE));
    }


    public void sendbtn(View view) {
        ServerInfo data = new ServerInfo();
        data.setArgs(argt.getText().toString());
        data.setCmd(cmdt.getText().toString());
        data.setTarget(trgt.getText().toString());
        send(data);
    }

    //Generic sender
    void send(ServerInfo data) {
        //Vibrate
        // Vibrate for 500 milliseconds

        if (swState) {
            new TcpSender().executeOnExecutor(AsyncTask.THREAD_POOL_EXECUTOR, data);
            new TcpSender().executeOnExecutor(AsyncTask.THREAD_POOL_EXECUTOR, data);
            new TcpSender().executeOnExecutor(AsyncTask.THREAD_POOL_EXECUTOR, data);
            new TcpSender().executeOnExecutor(AsyncTask.THREAD_POOL_EXECUTOR, data);
            new TcpSender().executeOnExecutor(AsyncTask.THREAD_POOL_EXECUTOR, data);
            new TcpSender().executeOnExecutor(AsyncTask.THREAD_POOL_EXECUTOR, data);
            new TcpSender().executeOnExecutor(AsyncTask.THREAD_POOL_EXECUTOR, data);
            new TcpSender().executeOnExecutor(AsyncTask.THREAD_POOL_EXECUTOR, data);
            new TcpSender().executeOnExecutor(AsyncTask.THREAD_POOL_EXECUTOR, data);
            new TcpSender().executeOnExecutor(AsyncTask.THREAD_POOL_EXECUTOR, data);
        } else {
            new TcpSender().executeOnExecutor(AsyncTask.THREAD_POOL_EXECUTOR, data);

        }
    }


    // Button functions
    public void msgHola(View view) {
        ServerInfo data = new ServerInfo();
        data.setArgs("Hola");
        data.setCmd("msg");
        send(data);
    }
    public void logout(View view) {
        ServerInfo data = new ServerInfo();
        data.setCmd("lo");
        send(data);
    }
    public void poweroff(View view) {
        ServerInfo data = new ServerInfo();
        data.setCmd("off");
        send(data);
    }
    public void killchrome(View view) {
        ServerInfo data = new ServerInfo();
        data.setArgs("chrome.exe");
        data.setCmd("kill");
        send(data);
    }
    public void killexplorer(View view) {
        ServerInfo data = new ServerInfo();
        data.setArgs("explorer.exe");
        data.setCmd("kill");
        send(data);
    }
    public void startexplorer(View view) {
        ServerInfo data = new ServerInfo();
        data.setArgs("explorer.exe");
        data.setCmd("web");
        send(data);
    }
    public void pornhub(View view) {
        ServerInfo data = new ServerInfo();
        data.setArgs("http://pornhub.com");
        data.setCmd("web");
        send(data);
    }
    public void redtube(View view) {
        ServerInfo data = new ServerInfo();
        data.setArgs("http://redtube.com");
        data.setCmd("web");
        send(data);
    }
    public void YtBtn(View view) {
        ServerInfo data = new ServerInfo();
        data.setArgs("http://www.youtube.com");
        data.setCmd("web");
        send(data);
    }
    public void franYtBtn(View view) {
        ServerInfo data = new ServerInfo();
        data.setArgs("http://www.youtube.com/channel/UCX6leuPdVe1hTdR6gGEzRfw");
        data.setCmd("web");
        send(data);
    }
    public void errorSuperGrave(View view) {
        ServerInfo data = new ServerInfo();
        data.setArgs("Error-super-grave.");
        data.setCmd("msg");
        send(data);
    }
    public void errorGraveBtn(View view) {
        ServerInfo data = new ServerInfo();
        data.setArgs("Error-grave.");
        data.setCmd("msg");
        send(data);
    }
    public void errorBtn(View view) {
        ServerInfo data = new ServerInfo();
        data.setArgs("Error.");
        data.setCmd("msg");
        send(data);
    }
    public void x10sw(View view) {
        swState = !swState;
    }



}
