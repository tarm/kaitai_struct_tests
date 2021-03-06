using NUnit.Framework;

namespace Kaitai
{
    [TestFixture]
    public class SpecMultipleUse : CommonSpec
    {
        [Test]
        public void TestMultipleUse()
        {
            var r = MultipleUse.FromFile(SourceFile("position_abs.bin"));
            Assert.AreEqual(r.T1.FirstUse.Value, 0x20);
            Assert.AreEqual(r.T2.SecondUse.Value, 0x20);
        }
    }
}
