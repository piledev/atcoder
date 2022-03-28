class PxUser:
    def __init__(self, name):
        self.name = name
    def say_hi(self):
        print("Hi! {0}".format(self.name))

class PxChildUser(PxUser):
    def __init__(self, name, age):
        super().__init__(name)
    def say_hello(self):
        print("Hello! {0}".format(self.name))

# モジュール単位でインポートされたタイミングで表示される?
# つまり、２回目の import では表示されない？
print("packageのmoduleをインポートしました")
