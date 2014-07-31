import java.lang.reflect.Constructor;

Class<Elvis> c = Elvis.class;
Constructor<Elvis> con = c.getDeclaredConstructor();
con.setAccessible(true);
Elvis instance = con.newInstance();
