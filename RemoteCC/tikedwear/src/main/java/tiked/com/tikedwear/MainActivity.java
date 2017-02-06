package tiked.com.tikedwear;

import android.app.Activity;
import android.content.Context;
import android.os.Bundle;
import android.support.annotation.NonNull;
import android.support.wearable.view.WatchViewStub;
import android.util.Log;
import android.view.View;

import com.google.android.gms.common.api.GoogleApiClient;
import com.google.android.gms.common.api.ResultCallback;
import com.google.android.gms.wearable.CapabilityApi;
import com.google.android.gms.wearable.MessageApi;
import com.google.android.gms.wearable.Node;
import com.google.android.gms.wearable.NodeApi;
import com.google.android.gms.wearable.Wearable;

import java.util.List;
import java.util.concurrent.TimeUnit;

public class MainActivity extends Activity {

    private GoogleApiClient mGoogleApiClient;
    Node mNode;
    String nodeId;
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        //mGoogleApiClient = new GoogleApiClient.Builder(this)
               // .addApi(Wearable.API)
                //.build();
        //resolveNode();
        //sendMessage("mensageeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeefdighffjghfjghkfgjkldfjgkñrejsgi9ehriughrekjhjgklefjgkhjrghjedfshgildfhlgiuhreughre");
        mGoogleApiClient = getGoogleApiClient(this);
        retrieveDeviceNode();
        sendToast();


    }

    private void resolveNode() {
        Wearable.NodeApi.getConnectedNodes(mGoogleApiClient)
                .setResultCallback(new ResultCallback<NodeApi.GetConnectedNodesResult>() {
                    @Override
                    public void onResult(@NonNull NodeApi.GetConnectedNodesResult connectedNodes) {
                        for(Node connectedNode : connectedNodes.getNodes()) {
                            mNode = connectedNode;
                            //Log.d(TAG,"Message Sender connected node"+mNode.getDisplayName());
                            sendMessage("Hello Hello");
                        }
                    }
                });
    }


    private void sendMessage(final String message) {
        if(mGoogleApiClient != null &&
                mGoogleApiClient.isConnected() &&
                mNode != null) {
            Log.d("okkkkkkk","Message is going to be sent to watch");

            Wearable.MessageApi.sendMessage(mGoogleApiClient,
                    mNode.getId(),
                    "/mobile_data",
                    message.getBytes())
                    .setResultCallback(new ResultCallback<MessageApi.SendMessageResult>() {
                        @Override
                        public void onResult(@NonNull MessageApi.SendMessageResult sendMessageResult) {
                            if(sendMessageResult.getStatus().isSuccess()) {
                                Log.e("okkkkkkkkkk","Message Succesfully sent to watch=>"+message);
                            } else {
                                Log.e("okkkkkkkkkk","Message FAILED TO BE SENT to watch=>"+message);
                            }
                        }
                    });
        }
    }

    public void buttonClick(View view) {
        sendMessage("sendddddddddddddddkdgjdfkgjfkgjfdkgjfdkjk");
    }

    private GoogleApiClient getGoogleApiClient(Context context) {
        return new GoogleApiClient.Builder(context)
                .addApi(Wearable.API)
                .build();
    }

    private void retrieveDeviceNode() {
        final GoogleApiClient client = getGoogleApiClient(this);
        new Thread(new Runnable() {
            @Override
            public void run() {
                client.blockingConnect(1000, TimeUnit.MILLISECONDS);
                NodeApi.GetConnectedNodesResult result =
                        Wearable.NodeApi.getConnectedNodes(client).await();
                List<Node> nodes = result.getNodes();
                if (nodes.size() > 0) {
                    nodeId = nodes.get(0).getId();
                }
                client.disconnect();
            }
        }).start();
    }
    private void sendToast() {
        final GoogleApiClient client = getGoogleApiClient(this);
        if (nodeId != null) {
            new Thread(new Runnable() {
                @Override
                public void run() {
                    client.blockingConnect(1000, TimeUnit.MILLISECONDS);
                    Wearable.MessageApi.sendMessage(client, nodeId, "holaaaa", null);
                    client.disconnect();
                }
            }).start();
        }
    }
}