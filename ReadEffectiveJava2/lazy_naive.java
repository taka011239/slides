public class Elvis {
    private static Elvis Instance;
    private Elvis() {...}
    public static Elvis getInstance() {
        if (Instance == null) {
            Instance = new Elvie();
        }
        return Instance;
    }

    public void leaveTheBuilding() {...}
}
