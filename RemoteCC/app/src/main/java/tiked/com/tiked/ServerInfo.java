package tiked.com.tiked;

import android.content.Context;
import android.os.Vibrator;

public class ServerInfo {
    static String target = "*";
    static String cmd = "";
    static String args = "";
    static Vibrator v;

    public static void setV(Vibrator v) {
        ServerInfo.v = v;
    }
    public void setArgs(String args) {
        ServerInfo.args = args.replaceAll("\\s+","");
    }

    public void setCmd(String cmd) {
        ServerInfo.cmd = cmd.replaceAll("\\s+","");
    }

    public void setTarget(String target) {
        String a = target.replaceAll("\\s+","");
        if (a == "all" || a == "todos") {a = "*";}
        ServerInfo.target = a;
    }
    public static String getCommand() {
        return cmd+" "+target+" "+args + "\n";
    }

}
