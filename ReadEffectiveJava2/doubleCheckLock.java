public class Elvis {
    private static volatile Elvis Instance;
    private Elvis() {...}
    public static Elvis getInstance() {
        if (Instance == null) {
            synchronized (Elvis.class) {
                if (Instance == null) {
                    Instance = new Elvis();
                }
            }
        }
        return Instance;
    }

    public void leaveTheBuilding() {...}
}
