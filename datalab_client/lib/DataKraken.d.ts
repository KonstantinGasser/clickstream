export declare class DataKraken {
    private API_WS;
    private URL_TIMEOUT_RATE;
    private URL_TIME;
    private CURRENT_URL;
    private LAST_CLICK;
    private BTN_DEFS;
    private WS_TICKET;
    private WEB_SOCKET;
    constructor(app_token: string);
    private sayHello;
    private open;
    private attach;
    private getReferrer;
    private getCampaign;
    private urlListener;
    private onClick;
    private onHover;
    private getDevice;
    private static Event;
    private static elapsed;
}
