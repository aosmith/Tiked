package tiked.com.tiked;

public class ServerInfo {
    static String target = "*";
    static String cmd = "";
    static String args = "";

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
