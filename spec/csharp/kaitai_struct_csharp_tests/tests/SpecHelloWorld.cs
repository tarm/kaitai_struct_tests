using System;

namespace Kaitai
{
    using NUnit.Framework;

    [TestFixture]
    public class SpecHelloWorld : CommonSpec
    {
        [Test]
        public void TestHelloWorld()
        {
            HelloWorld r = HelloWorld.FromFile(SourceFile("fixed_struct.bin"));
            Assert.AreEqual(r.One, 0x50);
        }
    }
}
