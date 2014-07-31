public class Elvis implements Serializable {
    private static final long SerialVersionUID = 6825273283542226860L;
    private static final Elvis INSTANCE = new Elvis();
    private Elvis() {...}
    public static Elvis getInstance() { return INSTANCE; }

    public void leaveTheBuilding() {...}

    private Object readResolve() {
        return INSTANCE;
    }
}
